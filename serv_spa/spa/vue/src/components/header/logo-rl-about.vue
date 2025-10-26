<template>
    <router-link to="/about" id="rl-about"><img src="../images/logo150x150.png" alt="logo"></router-link>
</template>

<style scoped>
#rl-about {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    padding: 1rem;
    height: 1rem;
    width: 1rem;
    box-shadow: 1px 2px 3px rgb(179, 210, 216);
}

#rl-about:hover {
    cursor: pointer;
    box-shadow: 0px 0px 6px rgb(179, 210, 216);
}

img {
    height: 2rem;
    border-radius: 50%;
    padding: 0.4rem;
}
</style>

<script>

</script>