<template>
    <nav>
        <v-snackbar v-model="snackbar" :timeout="4000" top color="primary success-text">
            <span>Dodano nową kartę do bazy!</span>
            <v-btn flat color="white" @click="snackbar = false">
                <v-icon>clear</v-icon>
            </v-btn>
        </v-snackbar>

        <v-snackbar v-model="snackbarLoginRequred" :timeout="4000" top color="primary success-text">
            <span>Wymagane logowanie!</span>
            <v-btn flat color="white" @click="snackbar = false">
                <v-icon>clear</v-icon>
            </v-btn>
        </v-snackbar>

        <v-snackbar v-model="snackbarLogout" :timeout="4000" top color="primary success-text">
            <span>Wylogowano</span>
            <v-btn flat color="white" @click="snackbarLogout = false">
                <v-icon>clear</v-icon>
            </v-btn>
        </v-snackbar>

        <v-snackbar v-model="snackbarLogin" :timeout="4000" top color="primary success-text">
            <span>Poprawnie zalogowano</span>
            <v-btn flat color="white" @click="snackbarLogin = false">
                <v-icon>clear</v-icon>
            </v-btn>
        </v-snackbar>

        <v-toolbar flat app>
            <v-toolbar-side-icon class="grey--text" @click="drawer = !drawer">
                <!-- burger menu icon, color zmieniony jak na ten tekstu -->
            </v-toolbar-side-icon>

            <v-toolbar-title class="text-uppercase grey--text">
                <span class="font-weight-light">Karty</span>
                <span> Edżentelmenów</span>
            </v-toolbar-title>
            <!-- tag <v-spacer> - to co nad nim bedzie do lewej 
            a to co po nim do prawej wcisniete -->
            <v-spacer></v-spacer> 

            <!-- dropdown menu -->
            <v-menu offset-y>
                <v-btn flat slot="activator" color="grey">
                    <v-icon left>expand_more</v-icon>
                    <span>menu</span>
                </v-btn>
                <v-list>
                    <v-list-tile v-for="link in links" :key="link.text" router :to="link.route">
                        <v-list-tile-title>{{ link.text }}</v-list-tile-title>
                    </v-list-tile>
                </v-list>
            </v-menu>

            <v-btn flat color="grey" @click="logout">
                <!-- wylogowanie -->
                <span>Wyloguj</span>
                <v-icon right>exit_to_app</v-icon>
            </v-btn>
        </v-toolbar>

    <!-- bez app nie dziala ladnie   -->
    <!-- na malych ekranach zaslania cale gorne menu, na duzych ladnie przsuwa kontent strony -->
        <v-navigation-drawer v-model="drawer" app class="primary">
            
            <v-layout column align-center>
                <v-flex class="mt-5">
                    <v-avatar size ="100" color="white" @click="drawer=false">
                        <img src="/cylinder-okragly.png">
                    </v-avatar>
                    <p class="white--text mt-3 subheading">Edżentelmeni</p>
                </v-flex>
                <v-flex class="mb-3">
                    <Popup @cardAdded="snackbar = true"  @LoginRequired="snackbarLoginRequred = true"/>
                </v-flex>
            </v-layout>

            <v-list>
                <v-list-tile v-for="link in links" :key="link.text" router :to="link.route">
                    <v-list-tile-action>
                        <v-icon class="white--text">{{ link.icon }}</v-icon>
                    </v-list-tile-action>
                    <v-list-tile-content>
                        <v-list-tile-title class="white--text">{{ link.text }}</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>

            </v-list>
        </v-navigation-drawer>
    </nav>
</template>

<script>
import axios from 'axios'
import Popup from './Popup'
export default {
    components: { Popup },
    data(){
        return{
            drawer: false, //niepokazywac na poczatku
            snackbar: false,
            snackbarLogout: false,
            isLogged: false,
            snackbarLogin: false,
            snackbarLoginRequred: false,
        }
    },
    computed:{
        links(){
            return this.$store.state.links;
        }
    },
    methods:{
        logout(){


            console.log("Procedura wylogowania wylogowania")
            var userLogin = localStorage.getItem('user')
            if (userLogin !== "" && userLogin !== "null"){
                console.log("Mozna wylogowac: ", userLogin)
                var myVue = this
                
                var myPromise = new Promise(function (resolve, reject){

                    var GoReturn = axios.post("http://127.0.0.1:3000/api/updateLogoutUserView", {"login": "a"} , //aktualizacja stanu uzytkownika w bazie mysql
                        { crossdomain: true })
                        .then(response=>{
                            console.log("Aktualizowanie stanu uzytkownika w mysql ", response.data); //printuje Tokena
                            return response.data
                        }).catch(error=>{ 
                            console.log("ERROR: ", error.response.data)
                            return null
                        });
                    if (GoReturn){
                        resolve(GoReturn)
                    }else{
                        reject("Wylogowanie Go błąd ")
                    }
                })

                myPromise.then(function(fromResolve){
                    console.log("Promise koniec: ", fromResolve)
                    myVue.isLogged = false

                }).catch(function(fromReject){
                    console.log("No jakis bld przy wylogowaniu", fromReject)
                    });

            }
            this.isLogged = false
            this.$router.push('/')



        
        },
        checkIsUserLogged(){
            console.log("CREATE: ", localStorage.getItem('token') )
            if (localStorage.getItem('token') === "null"){ //brak tokenu jest wylogowany
                this.isLogged = false
                console.log("NIEZALOGOWANY")
                return
            }
            if(localStorage.getItem('token') ===  null){//brak tokenu
                this.isLogged = false
                console.log("NIEZALOGOWANY")
                return
            }

        }
    },
    created(){
        console.log(this.isLogged)
        this.checkIsUserLogged()
    }
    
}
</script>
