<template>
    <div class="footer">
        <div class="navigation" v-on:click="goBack">
            <img src="../images/icons/arrow_back_24dp_E8EAED_FILL0_wght400_GRAD0_opsz24.svg" alt="back">
        </div>
        <router-link to="/" class="navigation">
            <img src="../images/icons/home_24dp_E8EAED_FILL0_wght400_GRAD0_opsz24.svg" alt="home">
        </router-link>
        <div class="navigation" v-on:click="goForward">
            <img src="../images/icons/arrow_forward_24dp_E8EAED_FILL0_wght400_GRAD0_opsz24.svg" alt="forward">
        </div>
    </div>
</template>


<script>
import MenuIcon from 'vue-material-design-icons/Menu.vue';
export default {
    components: {
        MenuIcon,
    },
    name: 'footerB',
    methods: {
        goBack() {
            this.$router.go(-1);
        },
        goForward() {
            this.$router.go(1);
        }
    }
};

</script>

<style scoped>
.footer {
    height: 3rem;
    width: 100%;
    max-width: 500px;
    display: flex;
    flex-direction: row;
    align-items: center;
    align-content: center;
    justify-content: space-between;
    padding: 0 2rem;
    margin: 0 0 0.5rem 0;
}

.navigation {
    height: 100%;
    width: 100%;
    margin: 0.3rem;
    border: 1px solid;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    border: 1px solid rgb(197, 214, 216);
    border-radius: 1rem;
    box-shadow: 1px 2px 3px rgb(128, 199, 211);
}

.navigation:hover {
    cursor: pointer;
    box-shadow: 0px 0px 7px rgb(128, 199, 211);
}

img {
    height: 100%;
}
</style>
