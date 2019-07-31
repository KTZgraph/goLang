<template>
  <div class="dashboard">
    <v-container fluid class="my-5">
      <v-layout row class="mb-3">
        <v-tooltip top>
          <v-btn small flat color="grey" @click="sortBy('question')" slot="activator">
            <v-icon left small>live_help</v-icon>
            <span class="caption text-lowercase">Karty pytania</span>
          </v-btn>
          <span>Sortuj według kart pytań</span>
        </v-tooltip>

        <v-tooltip top>
          <v-btn small flat color="grey" @click="sortBy('answer')" slot="activator">
            <v-icon left small>note</v-icon>
            <span class="caption text-lowercase">Karty Odpowiedzi</span>
          </v-btn>
          <span>Sortujwedług kart Odpowiedzi</span>
        </v-tooltip>
      </v-layout>


      <!-- pa-3 jest już nie potrzebne w <v-card> lepiej to wyglada jak sa koloery na samym brzegu-->
      <v-card  flat v-for="card in visibleCards" :key="card.due">
        <!-- dynamiczne budowanie stringa bedącego wartością klasy - składnia z $ i ciapki ukosne -->
        <v-layout row wrap :class="`pa-3 cardType ${card.type}`">
          
          <v-flex xs12 md6>
           <div class="caption grey--text">Treść</div>
           <div>{{card.text}}</div>
          </v-flex>
          
          <!-- <v-flex xs6 sm4 md2>
           <div class="caption grey--text">Dodana</div>
           <div>{{card.due}}</div>
          </v-flex> -->

          <v-flex xs6 sm4 md2>
            <div class="right">
              <v-chip small :class="`${card.type} caption my-2`">{{card.type}}</v-chip>
            </div>
          </v-flex>
        </v-layout>
        <!-- Dodanie po każdym elemncie cienkiej poziomej linii oddzielajacej komponenty -->
        <v-divider></v-divider>
      </v-card>
  
    </v-container>
  
  </div>
</template>

<script>
// import axios from 'axios'
import PouchDB from 'pouchdb'
export default {
  data(){
    return{
      cards: [],
      visibleCards: [],
      lastPropName: "",
      totalRows: null
    }
  },
  methods:{
    sortBy(propName){
        //sortowanie po typie karty
      if (propName){
        this.visibleCards = this.cards.filter( obj => {
          console.log(obj.type === "answer")
            if (obj.type === propName) {
              console.log(obj)
            return obj
            }
        })
      }//koniec Id
      else{
        this.visibleCards = this.cards
      }
    },
    daneZbazy(){
    // w linii komend jako admin na Windzie w folderze app 
    // C:\Users\dp\Desktop\PracowaniaProgramowania\frontend\app>add-cors-to-couchdb http://127.0.0.1:5984 -u root -p password  success    
    var objectVue = this;
    var db = new PouchDB('http://localhost:5984/golang_cards');
    var cardsTmp = []

    
      db.allDocs({include_docs : true}).then(function (res) {
        objectVue.totalRows = res.total_rows; //aktualizacja zmiennej majacej liczbe wierszy

        res.rows.forEach(function (entry){
          var tmp = {}
          tmp.due = entry.doc.Timestamp
          tmp.text = entry.doc.Text
          if (entry.doc.isquestion){
            tmp.type = 'question'
          }
          else{
            tmp.type='answer'
          }
          cardsTmp.push(tmp)
        })

      })

      this.cards = cardsTmp
      this.sortBy(this.lastPropName)
    },
    checkDiference(){
      var objectVue = this;
      //Sprawdza róznice w bazie danych - czy nie zostały dodane nowe dane
      var db = new PouchDB('http://localhost:5984/golang_cards');

      //tylko sprawdza liczbe wierszy, bez pobierania dokuemnto - ktore sa prototype tye
      var promiseRows = new Promise(function(resolve, reject) {
        //tu sie dzieje magia promisow
          var rowsCount = db.allDocs().then(function (res){
            return res.total_rows;
            })

            if (rowsCount){
              resolve(rowsCount)
            }
            else{
              reject("No nie chcial mi dac wynikow")
            }
      //koniec magii
      })

      //wywolanie magii
      promiseRows.then(function(fromResolve){
        if (fromResolve !== objectVue.totalRows){ //javascriptowe fantastyczne porownania
          //aktualizuj dane wyswietlane
          objectVue.daneZbazy()
          console.log("Nowe dane")
        }
      }).catch(function(fromReject){
        console.log(fromReject)
      })

    }
  },
  mounted() {
    // console.log("Pobranie wszsytkich kard z bazy")
    this.daneZbazy() //wczytuje na poczatku dane
        
    //cyklicznie ywkonujaca sie funkcja - do sprawdzania zmian w bazie co 10 s - żeby plynnie wyswietlac nowo dodane karty jesli zmiana zostanie zauwaona
      this.timer = setInterval(this.checkDiference, 10000) //na biezaco aktualizuje dane

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
  }
}


</script>

<style>
.cardType.question{
  border-left: 5px solid #212121;
}

.v-chip.question{
  background:#212121;
  color: white;
}

.cardType.answer{
  border-left: 5px solid #B0BEC5;
}

.v-chip.answer{
  background: #CFD8DC;
  color: black;

}

</style>

