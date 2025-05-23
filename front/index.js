console.log("Hello World")
console.log(x);
var x = 7;
var x = 10; //var can be redeclared
console.log("x:",x);

var fn = function display(){
    let a = 15;
    console.log("Display function called");
    console.log("a:",a);
}
fn()

//let & const in js
let a = 10;
let b = 20;
console.log(a)
//let a = 15; //not allowed

//console.log("student id:",studentId) //cannot access const and let before the initializtion -> Temporal deadzone
const studentId = 1;
console.log("student id:",studentId)

let n = 101;
console.log("m:",m) //should be undefined
//Block Scope
{
    //shadowing
    var m = 10; //global scope
    let n = 100; //block scope
    const o = 1000; //block scope
    console.log("n:",n) 
}
//console.log(n) //reference error 

//closures
function countClicks(){
    var count = 0;
    return function update(){
        count += 1;
        console.log("count:",count)
    }
}
const clickCounter = countClicks(); // only call once
clickCounter(); 
clickCounter();
clickCounter();

const clickCounter2 = countClicks();
clickCounter2();
clickCounter2();

//setTimeout + Closures
function doSomething(){
    for(let cn = 1;cn<=5;cn++){
    setTimeout(function(){console.log("doSomething:",cn);},cn*1000)
    }
}
doSomething();

//scoping
//Global scope: A variable accessible everywhere.
//Function scope: A variable accessible only within the function where it's declared.
//Block scope: A variable accessible only within { ... } blocks (like in if, for, etc.).
function test() {
 var isPrime = true;
}
//console.log(isPrime) //not defined
{
    var isPrime = true; //block scope but accessible anywhere
    let isEven = false; //block scope
}
console.log(isPrime); //works fine
//console.log(isEven); //not defined


//Callback
//Asynchronous
//map filter and reduce
//arrow function
//Destructuring
//async await
