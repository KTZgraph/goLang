
<template>
  <div class="cards">
    <v-container>
        <!-- widok gentelmana, ktory wybiera karty -->
            <v-layout column align-center>
                <!-- karta pytanie na srodku -->
                <v-flex class="mt-5">
                    <div class="cardQuestion">
                        <div class="card-main">
                            <v-avatar size ="100">
                                <img src="/moustache-white.png">
                            </v-avatar>
                            <p class = "pa-1 white--text text-md-center">{{ questionCard.text }}</p>
                        </div>
                    </div>
                </v-flex>
            </v-layout>
            <v-layout row wrap>
            <v-flex xs12 sm6 md4 lg2 v-for="card in visibleCards()" :key="card.name">
                <v-card flat class="text-xs-center ma-1 grey lighten-2 ">
                <v-responsive class="pt-4">
                    <v-card-text class="">

                <!-- karty odpowiedzi albo karty dane uzytkownika zlaeznie czy jest gentelmenem czy nie  -->
                    <div class="card" @click="chosenCard">
                        <div class="card-main">
                            <v-avatar size ="100" color="white mb-1">
                                <img src="/moustache-black.png">
                            </v-avatar>
                            <p class = "pa-1">{{ card.text }}</p>
                        </div>
                    </div>
                </v-card-text>
                </v-responsive>
                </v-card>

            </v-flex>
            </v-layout>

    </v-container>
  </div>
</template>

<script>
export default {
    data(){
      return{
          isLogged : false,
          gentelaman: false,
          cards: [
          { text: 'Te majce mniej niż 180cm wzrostu pocieszne istoty które nazywają siebie mężczyznami'},
          { text: 'Te majce mniej niż 180cm wzrostu pocieszne istoty które nazywają siebie mężczyznami'},
          { text: 'Te majce mniej niż 180cm wzrostu pocieszne istoty które nazywają siebie mężczyznami'},
          { text: 'Te majce mniej niż 180cm wzrostu pocieszne istoty które nazywają siebie mężczyznami'},
          { text: 'Te majce mniej niż 180cm wzrostu pocieszne istoty które nazywają siebie mężczyznami'},
          { text: 'Te majce mniej niż 180cm wzrostu pocieszne istoty które nazywają siebie mężczyznami'},
        ],
        
        userCards: [
          { text: 'karta uzyktkownika0'},
          { text: 'karta uzyktkownika1'},
          { text: 'karta uzyktkownika2'},
          { text: 'karta uzyktkownika3'},
          { text: 'karta uzyktkownika4'},
          { text: 'karta uzyktkownika5'},
          { text: 'karta uzyktkownika6'},
          { text: 'karta uzyktkownika7'},
          { text: 'karta uzyktkownika8'},
          { text: 'karta uzyktkownika9'}
        ],
        questionCard: { text: 'Pracuje 80 godzin tygodniowo i nadal nie stać mnie na ...'},
      }
    },
    methods:{
        chosenCard(){
            // console.log("Wybrano karte")
        },
        visibleCards(){
            // console.log("Visible cards function")
            if (this.gentelaman===true){
                return this.cards
            }
            else{
                return this.userCards
            }
        },
        getUserCard(){
            //pobieranie z bazy kard wylosowanych dla danego uzytkownika
            localStorage.getItem('user')
        }
    },
    created(){
        //Zbuduje sie tylko i wylacznie wtedy gdy uzytkownik jest zalogowany, w przeciwnym wypadku przekierowuj do zalogowania

        if (localStorage.getItem('token') === "null"){
            this.isLogged = true
            console.log("TUTAJ: ", localStorage.getItem('token') === null)
            this.$router.push('/') 

        }
        if(localStorage.getItem('token') ===  null){
            this.$router.push('/') 

        }
        console.log("LOCAL STOARAGE: ", localStorage.getItem('token'))
        console.log("LOCAL STOARAGE User: ", localStorage.getItem('user'))

        console.log("Wszyscy uzytkownicy: ", this.$store.getters.allUsers)
        this.getUserCard()

    }
}
</script>


<style>
.card {
  width: 200px;                 /* Set width of cards */
  height: 300px;
  border: 5px solid #212121;    /* Set up Border */
  border-radius: 20px;           /* Slightly Curve edges */
  overflow: hidden;             /* Fixes the corners

  display: flex;                /* Children use Flexbox */
  flex-direction: column;       /* Rotate Axis */
  background: white;
}

.cardQuestion{
  background: #212121;
  width: 200px;                 /* Set width of cards */
  height: 300px;
  border: 5px solid black;    /* Set up Border */
  border-radius: 20px;           /* Slightly Curve edges */
  overflow: hidden;             /* Fixes the corners

  display: flex;                /* Children use Flexbox */
  flex-direction: column;       /* Rotate Axis */
}


.card-main {
  display: flex;              /* Children use Flexbox */
  flex-direction: column;     /* Rotate Axis to Vertical */
  justify-content: center;    /* Group Children in Center */
  align-items: center;        /* Group Children in Center (+axis) */
  padding: 15px 0;            /* Add padding to the top/bottom */
}

.material-icons {
  font-size: 36px;
  color: #D32F2F;
  margin-bottom: 5px;
}


.main-description {
  color: #D32F2F;
  font-size: 12px;
  text-align: center;
}
</style>
