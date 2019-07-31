<template>
  <div class="team">
    <h1 class="subheading grey--text">Wszyscy Dżentelmani</h1>

    <v-container class="my-5">
        <v-layout row wrap>
          <v-flex xs12 sm6 md4 lg3 v-for="person in allUsers" :key="person.login">
            <v-card flat class="text-xs-center ma-3">
              <v-responsive class="pt-4">
                  <v-avatar size="100" class="grey lighten-2">
                    <img src="/user-cylinder.png">
                  </v-avatar>
              </v-responsive>
              <v-card-text>
                <div  class="subheading">{{ person.login }}</div>

                <div v-if="person.userDescription" class="grey--text">{{ person.userDescription }}</div>
                <div v-else class="subheading white--text">_</div>

              </v-card-text>
              <v-card-actions>
                <v-btn flat color="grey">
                  <v-icon small left>grade</v-icon>
                  <span>{{ person.points }}</span>
                </v-btn>
              </v-card-actions>
            </v-card>

          </v-flex>
        </v-layout>
    </v-container>
  </div>
</template>

<script>
import  axios from 'axios'
export default {
  data(){
    return{
      allUsers: []
    }
  },
  methods:{
    getAllUsers(){
      var objectVue = this;

      var promiseUsers = new Promise(function(resolve, reject) {
        // var allUsersTmp = axios.get("http://127.0.0.1:3000/api/getUsersView" ,{ crossdomain: true }
        // )

        var allUsersTmp = axios({
          method: 'get',
          url: 'http://127.0.0.1:3000/api/getUsersView',
          params: {'Token': "TokenZJs"}
        })
          .then(function(res){
            console.log("--------------------Users response:------------------------", res.data.allUsers)
              return res.data.allUsers //zwraca uzytkownikow
          }).then(function(res){
            objectVue.team = res
            return res
          })

        if(allUsersTmp){
          resolve(allUsersTmp)
        }
        else{
          reject("Nie udalo sie pobrac uzytkownikow")
        }

      })

    //Promise.all czeka na wykonanie wszystkich
     Promise.race([promiseUsers]).then(function(){
        promiseUsers.then(function(fromResolve){
          objectVue.allUsers = fromResolve
          return fromResolve
        })
      })

    }
  },
  mounted(){
    this.getAllUsers()
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

