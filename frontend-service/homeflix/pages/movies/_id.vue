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
                        <button class="button">{{torrent.type}} {{torrent.quality}}</button>
                    </span>
                </p>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import Loading from '../../components/loading.vue'
export default {
    name: "single-movie",
    data() {
        return {
            movie: null,
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
        }
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
</style>