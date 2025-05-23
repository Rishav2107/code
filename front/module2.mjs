//.mjs tells Node.js or the browser:
//“This file uses import/export syntax.”

// import name from './module1.mjs' //default import 
import {n2,n3} from './module1.mjs' //for non-default we need to pass the exact names that was exported

//console.log(name) //default
console.log(n2)
console.log(n3)