<template>
<div>
    <Loading v-if="isProcessing"/>
    <div class="homeflix">
        <img src="../assets/img/background.jpg" alt="">
        <div class="text-container">
            <div class="text">
                <span class="mini-heading">Homeflix</span>
                <h1><span>Home</span>Flix</h1>
                <a href="#movie-grid" class="button">View Movies</a>
                <button class="button" @click="UpdateCollection">Update Collection</button>
            </div>
        </div>
    </div>
</div>
    
</template>

<script>
import axios from 'axios'
    export default{
        name:"homeflix",
        data(){
            return{
                isProcessing:false
            }
        },
        methods : {
            UpdateCollection (){
                this.isProcessing = true
                const data = axios.get('http://localhost:9090/updatemovie')
                setTimeout(function(){
                    this.isProcessing = false
                    location.reload()
                },5000)
            },
        }
    }
</script>

<style lang="scss" scoped>
    .homeflix {
        height: 400px;
        position: relative;
        @media (min-width: 750px) {
            height: 500px;
        }
        &::after {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            height: 100%;
            width: 100%;
            background-color: rgba(0,0,0,0.6);
        }
        img {
            width: 100%;
            height: 100%;
            object-fit: cover;
        }
        .text-container {
            z-index: 99;
            position: absolute;
            top: 0;
            margin: 0;
            width: 100%;
            height: 100%;
            display: flex;
            flex-direction: columns;
            justify-content: center;
            .text {
                padding: 100px 16px;
                width: 100%;
                max-width: 1400px;
                margin: 0 auto;
            }
            .mini-heading {
                font-weight: 600;
                font-size: 18px;
                text-transform: uppercase;
                color: #ffffff;
                margin-bottom: 8px;
                @media (min-width: 750px) {
                    font-size: 22px;
                }
            }
            h1 {
                color: #fff;
                font-size: 64px;
                font-weight: 200;
                margin-bottom: 8px;
                @media (min-width: 750px) {
                    font-size: 84px;
                }
                span {
                    font-weight: 500;
                }
            }
            .button {
                font-size: 20px;
                align-self: flex-start;
            }
        }
    }
</style>