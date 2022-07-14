<template>
    <div class="loading">
        <span></span>
        <h1>Please Wait...</h1>
    </div>
</template>

<script>
export default{}
</script>

<style lang="scss" scoped>
.loading {
    height: 100vh;
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    z-index: 1000;

    @keyframes circle-loading {
        to {
            transform: rotateZ(360deg);
        }
    }

    span {
        display: block;
        width: 70px;
        height: 70px;
        border-radius: 50%;
        border: 2px solid transparent;
        border-top-color: #fff;
        animation: circle-loading 1200ms ease infinite;
    }

    h1 {
        color: #fff;
    }
}
</style>