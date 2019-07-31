import Vue from 'vue';
import Vuex from 'vuex';
// import axios from 'axios';

Vue.use(Vuex);

export const store = new Vuex.Store({
    state: {
        links:[
            { icon: 'dashboard', text: 'Gra', route: '/game'},
            { icon: 'person', text: 'Dżentelmani', route: '/dzentelmani'},
            { icon: "folder", text: 'Wszystkie Karty', route: '/karty'},
            // { icon: 'exit_to_app', text: 'Wyloguj', route: '/logout'},
            // { icon: 'get_app', text: 'Zaloguj', route: '/login'},
            // { icon: 'account_circle', text: 'Rejestracja', route: '/register'},
        ],
        allLoggedUsers:[]
    },
    getters:{
        allUsers: state => {
            return state.allLoggedUsers
        }
    },
    mutations:{
        //usuwanie po wylogowaniu
        removeLoggedUser: (state, payloadUser)=> {
            var elementId = state.allLoggedUsers.indexOf(payloadUser) //znajdowanie id elementu do usuniecia
            delete state.allLoggedUsers[elementId] //usuwanie kontkretnego loginu

        },
        //dodawanie nowego zalogowanego uzytkownika
        addLoggedUser: (state, payloadUser)=> {
            state.allLoggedUsers.push(payloadUser)

        }
    },
    actions:{
        addLoggedUser: (context, payloadUser) =>{
            context.commit('addLoggedUser', payloadUser)
        },
        removeLoggedUser: (context, payloadUser) =>{
            context.commit('removeLoggedUser', payloadUser)
        }
    }
})
