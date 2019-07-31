<template>
<div id="login">
      <!-- okno  jesti cos ma sie wyswietlic-->
      <v-snackbar v-model="snackbar" :timeout="4000" top color="red">
        <span>Nie można zalogować</span>
        <v-btn flat color="white" @click="snackbar = false">
            <v-icon>clear</v-icon>
        </v-btn>
      </v-snackbar>

    <!-- Okno rejestracji -->
    <v-container class="my-5">
        <v-layout column align-center>
            <v-flex id ="flex" xs12 sm6 md4 lg3 class="primary">
                <v-card flat class="text-xs-center ma-3" min-width="600px">
                    <v-avatar size ="100" color="white mt-3">
                        <img src="/cylinder-okragly.png">
                    </v-avatar>
                    <p class="primary--text mt-3 headline">Edżentelmeni</p>
                    <v-card-text>
                        <v-form class="px-3" ref="form">
                          <v-text-field :rules="loginRules" label="login" v-model="userLogin" prepend-icon="person"></v-text-field>
                          <v-text-field :rules="loginRules" label="hasło" v-model="userPassword"  :type="'password'"  prepend-icon="vpn_key"></v-text-field>
                        </v-form>
                    </v-card-text>
                    <!-- <v-card-actions> -->

                          <v-btn align-center justify-center class="primary text-md-center" @click="login">
                              Zaloguj
                          </v-btn>
                    <!-- </v-card-actions> -->
                    <v-card-text>
                    <v-list-tile router :to="'/register'">
                        <v-list-tile-title class="text">
                          <p class="text-md-center">
                            Nie masz konta? Zarejestruj się!
                          </p>
                        </v-list-tile-title>
                    </v-list-tile>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</div>
</template>

<script>
import axios from 'axios'

export default {
    data(){
        return{
          snackbar: false,
            userLogin : '',
            userPassword: '',
            loginRules: [
              v => v.length > 0 || 'Pole nie może być puste'
            ] //nie sprawdzam hasla, bo to przy rejestracji
        }
    },
    methods:{
        login(){
          if (this.$refs.form.validate()){ //valinnaj JS

            axios.post("http://127.0.0.1:3000/api/loginUserView", 
             {"login": this.userLogin, "password": this.userPassword}, //DANE DO LOGOWANIA
              { crossdomain: true })
              .then(response=>{
                  console.log("Response data: ", response.data); //printuje Tokena
                  
                  if (response.data){
                    console.log("Dostał tokena")
                    console.log("DANE Z VUEXA DLA USERA")
                    localStorage.setItem('token',response.data);
                    localStorage.setItem('user', JSON.stringify(this.userLogin));

                     
                    var local = localStorage.getItem('token');
                    console.log("______________________________TOKEN_____________________________")
                    console.log(localStorage.getItem('token'))
                    var localU = localStorage.getItem('user');

                    //dodanie loginu do listy wszystkich zalogowanych użytkowników
                    this.$store.dispatch('addLoggedUser', localStorage.getItem('user'))
                    
                    
                    axios.post("http://127.0.0.1:3000/api/updateLoggedUserView", {"login": this.userLogin} , //aktualizacja stanu uzytkownika w bazie mysql
                    { crossdomain: true })
                    .then(response=>{
                        console.log("Aktualizowanie stanu uzytkownika w mysql ", response.data); //printuje Tokena
                    }).catch(error=>{ 
                        console.log("ERROR: ", error.response.data)
                    });

                    console.log("Wszyscy uzytkownicy: ", this.$store.getters.allUsers)

                    console.log("LOCAL STOARAGE: ", local)
                    console.log("LOCAL STOARAGE User: ", localU)

                    this.$router.push('game') //przekierowanie dla one sigle page application  window.location.replace("http://localhost:8080/game");

                  }
                  else{ //coś poszło nie tak
                    console.log("Zły lolgin")
                    this.snackbar = true
                    this.userLogin = ''
                    this.userPassword = ''
                    localStorage.setItem('token',null);
                    localStorage.setItem('user', null);
                  }
              })  
              .catch(error=>{ 
                console.log("ERROR: ", error.response.data)
              });

          }else{ //dlugi if
            this.userLogin = ''
            this.userPassword = ''
          }

        }
    }
}
</script>

<style>
.v-card{
      border-radius: 3px;
}

#flex{
    border-radius: 5px;
}
</style>
