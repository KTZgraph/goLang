<template>
    <h1>Wylogowanie</h1>
</template>

<script>
export default {
    data(){
        return{

        }
    },
    methods:{
        logout() {
        // remove user from local storage to log user out
            localStorage.removeItem('user');
        }
    }
}

</script>
