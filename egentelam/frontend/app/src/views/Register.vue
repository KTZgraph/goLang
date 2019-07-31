<template>
<div id="register">
    <!-- okno  jesti cos ma sie wyswietlic-->
            <v-snackbar v-model="snackbar" :timeout="4000" top :color="`${snackbarColour}`">
            <span>{{snackbarText}}</span>
            <v-btn flat color="white" @click="snackbar = false">
                <v-icon>clear</v-icon>
            </v-btn>
        </v-snackbar>
    <!-- Rejestracja -->
    <v-container class="my-5">
        <v-layout column align-center>
            <v-flex id="flex" xs12 sm6 md4 lg3 class="primary">
                <v-card flat   class="text-xs-center ma-3" min-width="600px">
                    <v-avatar size ="100" color="white mt-3">
                        <img src="/cylinder-okragly.png">
                    </v-avatar>
                    <p class="primary--text mt-3 headline">Edżentelmeni</p>
                    <v-card-text>

                        <v-form class="px-3" ref="form">
                            <v-text-field :rules="rulesInput" label="login" v-model="userLogin" prepend-icon="person"></v-text-field>
                            <v-text-field :rules="rulesPassword" label="hasło" v-model="userPassword" :type="'password'"  prepend-icon="vpn_key"></v-text-field>
                            <v-text-field :rules="rulesPassword" label="potwierdź hasło " v-model="userPassword2"  :type="'password'"  prepend-icon="vpn_key"></v-text-field>
                            <v-text-field  label="opis nieobowiązkowy " v-model="userDescription" prepend-icon="favorite"></v-text-field>
                        </v-form>

                    </v-card-text>
                    <!-- <v-card-actions> -->
                        <v-btn class="primary" @click="register">
                            Zarejestruj
                        </v-btn>
                    <!-- </v-card-actions> -->
                <v-card-text>
                    <v-list-tile router :to="'/'">
                            <v-list-tile-title class="mt-3 subheading" >
                                <p class="text-md-center">
                                    Masz już konto? Zaloguj się! 
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
            snackbarColour: 'red',
            snackbarText:  '',
            userLogin : '',
            userPassword: '',
            userPassword2: '',
            userDescription: '',
            // rulesPassword: [ v => v.match(/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z]{8,}$/) || 'Nieprawidłowe hasło'],
            rulesPassword: [v => v.length > 0 || 'Hasło nie może być puste'],
            rulesInput: [v => v.length > 0 || 'Pole nie może być puste']
        }
    },
    methods:{
        register(){
            //dodwanie nowego uzytkownika do bazy
            console.log("Rejestracja uzytkownika")
            if (this.$refs.form.validate()){ //valinnaj JS
                console.log("WALIDAJC AOK")
                
                if (this.userPassword != this.userPassword2){
                    this.rulesPassword= [false || 'Podane hasla różnią się']
                    console.log("Hasla sei roznia")
                    this.userPassword = ''
                    this.userPassword2 = ''
                    return
                }
                if(this.userPassword === this.userPassword2){
                    //hasla ok, login ok - to rejestracja w bazie
                    this.rulesPassword = [true || 'Rejestracj']
                    console.log("OK register")
                    
                    //dodawanie do bazy
                    axios.post("http://127.0.0.1:3000/api/addNewUserView" ,
                        {"login": this.userLogin, "password": this.userPassword, "userDescription": this.userDescription}, 
                        { crossdomain: true })
                        .then(response=>{
                            console.log(response.data);
                            if (response.data === "[addNewUserView] Dodano uzytkownika do bazy"){
                                this.snackbar = true
                                this.snackbarColour = 'green darken-4 white--text mt-4'
                                this.snackbarText= "Utworzono konto :)"
                                // window.location.replace("http://localhost:8080/");
                            }
                            if (response.data === "Login zajęty"){
                                this.snackbar = true
                                this.snackbarColour = "red"
                                this.snackbarText= "Login zajety :<"
                            }
                            // else{
                            //     //Zły login albo hasło
                            //     this.loginRules= [false || 'Nieprawidłowe dane'] 
                            // }


                        })  

                        .catch(error=>{ 
                            console.log("ERROR: ", error.response.data)
                        });
                    



                    this.userLogin = ''
                    this.userPassword = ''
                    this.userPassword2 = ''
                    return 

                }
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

