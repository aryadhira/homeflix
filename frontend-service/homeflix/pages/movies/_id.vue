<template>
    <Loading v-if="$fetchState.pending"/>
    <div v-else class="container single-movie">
        <NuxtLink class="button" :to="{name:'index'}">Back</NuxtLink>
        <div class="movie-info">
            <div class="movie-img">
                <img :src="`${movie.large_cover_image}`" alt=""/> 
            </div>
            <div class="movie-content">
                <h1>{{movie.title}}</h1>
                <p class="movie-fact"><span>Released Year : </span>{{movie.year}}</p>
                <p class="movie-fact"><span>Genre : </span>
                    <span v-for="(genre,index) in movie.genres" :key="index">
                        {{genre}} ,
                    </span>
                </p>
                <p class="movie-fact"><span>Runtime : </span>{{movie.runtime}} minutes</p>
                <p class="movie-fact"><span>Rating : </span>{{movie.rating}}</p>
                <!-- next feature <p class="movie-fact"><span>Trailer : </span>https://youtube.com/watch?v={{movie.yttrailercode}}</p> -->
                <p class="movie-fact tagline"><span>Synopsis : </span>{{movie.synopsis}}</p>
                <p class="movie-fact"><span>Play Now : </span>
                    <span v-for="(torrent,index) in movie.torrents" :key="index">
                        <button class="button" @click="popModal(torrent.type,torrent.quality)">{{torrent.type}} {{torrent.quality}}</button>
                    </span>
                </p>
            </div>
        </div>
        <div v-if="showModal" class="modal">
            <!-- Modal content -->
            <div class="modal-content">
                <div class="modal-header">
                    <span class="close" @click="showModal = false">&times;</span>
                    <h1>{{movie.title}} {{type}} {{quality}}</h1>
                </div>
                <div class="modal-body">
                    <p>Some text in the Modal..</p>
                </div>
                <div class="modal-footer">
                    <button class="button" @click="showModal = false">Close</button>
                </div>
                
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import Loading from '../../components/loading.vue'
const WebTorrent = require('webtorrent')
export default {
    name: "single-movie",
    data() {
        return {
            movie: null,
            showModal: false,
            quality: "",
            type: "",
        };
    },
    async fetch() {
        this.getDetailMovie();
    },
    fetchDelay: 1000,
    methods: {
        async getDetailMovie() {
            const id = this.$route.params.id;
            const url = "http://localhost:9090/getmoviebyid/" + id;
            const data = axios.get(url);
            const result = await data;
            this.movie = result.data.Data;
        },
        async popModal(type,quality) {
            this.showModal = true
            this.quality = quality
            this.type = type
            let hash = this.getHash()
            console.log("hash",hash,this.movie.title_long,this.quality)

            const url = "http://localhost:9090/getmagneturl?title="+this.movie.title_long+"&qlty="+this.quality+"&hash="+hash

            const data = axios.post(url)
            const res = await data
            const magneturl = res.data.magnet
            console.log(res.data.magnet)

            this.streamMovie(magneturl)
        },
        getHash() {
            let torrents = this.movie.torrents
            let hash = ""
            for (var i=0; i<torrents.length; i++) {
                if (torrents[i].type == this.type && torrents[i].quality == this.quality){
                    hash = torrents[i].hash
                }
            }
            return hash
        },
        streamMovie(magnet) {
            const client = new WebTorrent()

            client.on('error',function(err){
                console.log(err.message)
            })

            console.log("adding torrent")
            client.add(magnet,this.onTorrent)

        },
        onTorrent(torrent){
            console.log(torrent.name)
            console.log(torrent.magnetURI)
            const file = torrent.files.find(function (file) {
                return file.name.endsWith('.mp4')
            })
            file.appendTo('.modal-body')
        },
    },
    components: { Loading }
}
</script>

<style lang="scss" scoped>
.single-movie{
    color: #fff;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 32px 16px;

    .button{
        align-self: flex-start;
        margin-bottom: 32px;
    }

    .movie-info{
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 32px;
        color: #fff;
        @media (min-width: 800px) {
            flex-direction: row;
            align-items: flex-start;
        }
        .movie-img{
            img{
                max-height: 500px;
                width: 100%;

                @media (min-width: 800px) {
                    max-height: 700px;
                    width: initial;
                }
            }
        }
        .movie-content{
            h1{
                font-size: 56px;
                font-weight: 400;
            }
            .movie-fact{
                margin-top: 12px;
                font-size: 20px;
                line-height: 1.5;

                button{
                    margin-left: 5px;
                }
            }
            span{
                font-weight: 600;
            }
        }
        .tagline {
            font-style: italic;
            span{
                font-style: normal;
            }
        }
    }
}

.modal {
  display: flex; /* Hidden by default */
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0,0,0); /* Fallback color */
  background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
  color: #fff;

    .modal-content {
    background-color: #0f0606;
    margin: auto;
    padding: 20px;
    border: 1px solid rgb(12, 10, 10);
    width: 90%;

        .modal-header {
            padding-bottom: 20px;
        }
        .modal-body {
            height: 600px;
            padding-bottom: 20px;
        }
        .modal-footer {
            display: flex;
            flex-direction: row-reverse;
        }
    }

    .close {
    color: #aaaaaa;
    font-size: 28px;
    font-weight: bold;
    float: right;
    }

    .close:hover,
    .close:focus {
    color: #000;
    text-decoration: none;
    cursor: pointer;
    }
}


</style>