<template>
    <v-dialog max-width="600px" v-model="dialog">
        <!-- to widac na przycisku -->
        <v-btn flat slot="activator" class="success primary--text" >Dodaj Nową Kartę</v-btn>
        <v-card>
            <v-card-title>
                <!-- to widac w wyswietlanym oknie -->
                <h2>Nowa karta</h2>
            </v-card-title>
            <v-card-text>
                <v-form class="px-3" ref="form">

                <v-select
                  :items="['pytanie', 'odpowiedź']"
                  label="Typ karty*"
                  required
                  prepend-icon= folder
                  v-model="cardType"
                  :rules="cardTypeRules"
                ></v-select>
                    <v-textarea label="Treść*" v-model="cardText" prepend-icon="edit" :rules="inputRules"></v-textarea>
                    <v-spacer></v-spacer>

                    <v-btn class="green darken-4 white--text mt-4" @click="submit" :loading="loading" prepend-icon="edit" >
                        <v-icon >done</v-icon>
                    </v-btn>
                    </v-form>
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script>
import format from 'date-fns/format'
import axios from 'axios'

export default {
    data(){
        return{
            cardText: '',
            cardType: '',
            due: null,
            inputRules: [
                v => v.length >= 2 || 'Treść karty musi mieć minimum 2 znaki'
            ],//array of diferent rules
            cardTypeRules: [ v=> v.length > 0 || 'Typ karty jest wymagany' ],
            loading: false,
            dialog: false
        }
    },
    methods:{
        submit(){
            //pobiera dane i wysyla na backend jsona dla karty
            /*{	"id" : 1,        
	            "IsQuestion" : true ,
	            "Blank" :1     ,
                "texst": "ALA MA KOTA" 
            }*/
            if (localStorage.getItem('token') === "null"){
                this.isLogged = true
                console.log("TUTAJ: ", localStorage.getItem('token') === null)
                this.$router.push('/')
                this.$emit('LoginRequired')
                return
            }
            if(localStorage.getItem('token') ===  null){
                console.log("TUTAJ: ", localStorage.getItem('token') === null)
                this.$router.push('/')
                this.$emit('LoginRequired')
                return
            }
            else{ //Dodawanie tylko dla zalogowanych
                console.log("DODAWANIE KARTY")

                if (this.$refs.form.validate()){ // valinnaj JS
                    var blank = 0
                    if (this.cardText.match(/_/g)){
                        blank = this.cardText.match(/_/g).length;
                    }
                    if (this.cardType === "pytanie" && blank === 0){
                        //walidacja - karta pytanie musi miec min 1 pole puste
                        this.inputRules = [false || 'Karta pytanie musi miec min jedno pole na odpwiedz (znak podłogi _)'] //info zwrotne w dialogu dla uzytkownika
                    }

                    var cardObj = {}
                    if (this.cardType == "pytanie"){
                        cardObj.isQuestion = true
                    }else{
                        cardObj.isQuestion = false
                    }
                    cardObj.blank = blank
                    cardObj.text = this.cardText

                    axios.post("http://127.0.0.1:3000/api/addNewCardView" ,cardObj, { crossdomain: true })
                    .then(response=>{
                            console.log(response);
                        })  
                        .catch(error=>{ //javascript ma swoje posrane zwrotki
                        //400 bad request wychodzi z asynchronicznosci tutaj, i wszyscy tak to obsluguja
                            //Sprawdzanie czy to na pewno blad czy po prostu javascript sie fąfla -,-
                            if ( error.response.data == "[addNewCardView] Dodano nową kartę"){
                                this.$emit('cardAdded')
                                return
                            }
                            if(error.response.data == "[updateUserPointsView][Error] Nie podano 'text'"){
                                // To sie nie zadzieje, bo js ma walidacje
                                console.log("Nie podano tekstu karty")
                                return
                            }
                            else{//tu jest blad
                                console.log("ERROR: ", error.response.data)
                                return
                            }
                        });
                
                } ///dlugi if
            }
        }
    },
    computed:{
        formattedDate(){
            return this.due ? format(this.due, 'Do MMM YYYY') : ''
        }
    }
}
</script>
