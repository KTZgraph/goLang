//connector couchdb
//Inaczej pobieranie kard, które są dostępne publicznie będzie nieodświeżało widoku wszystkich kart po dodaniu przez backend danych 
// gorsze rozwiazanie to zasluchiwanie ciagle na zmiany w bazie i wtedy wykonanie w widoku pobrania jeszcze raz danych, ale komunikacja źre za dużo, więc nie

// const NodeCouchDb = require('node-couchdb');

// // node-couchdb instance with default options
// const couch = new NodeCouchDb();
 
// // node-couchdb instance with Memcached
// const MemcacheNode = require('node-couchdb-plugin-memcached');
// const couchWithMemcache = new NodeCouchDb({
//     cache: new MemcacheNode
// });
 
// // node-couchdb instance talking to external service
// const couchExternal = new NodeCouchDb({
//     host: 'couchdb.external.service',
//     protocol: 'https',
//     port: 6984
// });
 
// // not admin party
// const couchAuth = new NodeCouchDb({
//     auth: {
//         user: 'root',
//         pass: 'password'
//     }
// });


// export default couch;

