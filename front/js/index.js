//PDF Notes: https://github.com/alok722/namaste-javascript-notes
//Javascript works on Two phase execution
//phase1 : Memory assignment
//phase2 : Code execution
console.log("Hello World");
console.log(x);
var x = 7;
var x = 10; //var can be redeclared
console.log("x:", x);

var fn = function display() {
  let a = 15;
  console.log("Display function called");
  console.log("a:", a);
};
fn();

//let & const in js
let a = 10;
let b = 20;
console.log(a);
//let a = 15; //not allowed

//console.log("student id:",studentId) //cannot access const and let before the initializtion -> Temporal deadzone
const studentId = 1;
console.log("student id:", studentId);

let n = 101;
console.log("m:", m); //should be undefined
//Block Scope
{
  //shadowing
  var m = 10; //global scope
  let n = 100; //block scope
  const o = 1000; //block scope
  console.log("n:", n);
}
//console.log(n) //reference error

//closures
function countClicks() {
  var count = 0;
  return function update() {
    count += 1;
    console.log("count:", count);
  };
}
const clickCounter = countClicks(); // only call once
clickCounter();
clickCounter();
clickCounter();

const clickCounter2 = countClicks();
clickCounter2();
clickCounter2();

//setTimeout + Closures
function doSomething() {
  for (let cn = 1; cn <= 5; cn++) {
    setTimeout(function () {
      console.log("doSomething:", cn);
    }, cn * 1000);
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

//objects
let obj = {
  first: "Rishav",
  second: "Kumar",
  greet: function displayGreet() {
    console.log("Hello World");
  },
};
console.log(obj.first, " ", obj.second);
obj.greet();

//arrays
let arr = [34, 21, 23];

//objects
const employee = { company: "google", city: "hyderbad" };
console.log("employee object:", employee);

//arrow functions
const hello = (value) => {
  console.log("hello ", value);
};
hello("Rishav");

// const ex = {
//     name: "Rishav",
//     role: "senior sde",
//     show: function(){
//         function print(){
//             console.log("ex: ",this.name,"is a",this.role);
//         }
//         print();
//     }
// }
// ex.show()  //o/p ex:  undefined is a undefined

//arrow functions capture this lexically, so this.name and this.role refer to the properties in ex.
const ex = {
  name: "Rishav",
  role: "senior sde",
  show: function () {
    return () => {
      console.log("ex: ", this.name, "is a", this.role);
    };
  },
};
ex.show()();

//global window
console.log(this);
console.log("window : x ",this.x);

//Callback
//Asynchronous
//map filter and reduce
//Destructuring
//async await
//promises


/* O/P
[Running] node "/Users/rishavkumar/Documents/code/front/js/index.js"
Hello World
undefined
x: 10
Display function called
a: 15
10
student id: 1
m: undefined
n: 100
count: 1
count: 2
count: 3
count: 1
count: 2
true
Rishav   Kumar
Hello World
employee object: { company: 'google', city: 'hyderbad' }
hello  Rishav
ex:  Rishav is a senior sde
Window {window: Window, self: Window, document: document, name: '', location: Location, …}
window : x  10
doSomething: 1
doSomething: 2
doSomething: 3
doSomething: 4
doSomething: 5
*/