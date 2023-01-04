/*const loginForm = document.querySelector("#login-form");
const loginInput = loginForm.querySelector("input")

const greeting = document.querySelector("#greeting")

const HIDDEN_CLASSNAME = "hidden"
const USERNAME_KEY = "username"

function onLoginSubmit(event) {
    event.preventDefault();
    const usernameThatUserWrote = loginInput.value;
    loginForm.classList.add(HIDDEN_CLASSNAME);
    localStorage.setItem(USERNAME_KEY, usernameThatUserWrote);
    paintGreeting();
}

function paintGreeting(){
    const username = localStorage.getItem(USERNAME_KEY);
    greeting.innerText = `Hello ${username}`;
    greeting.classList.remove(HIDDEN_CLASSNAME);
}


const saveUsername = localStorage.getItem(USERNAME_KEY);

if(saveUsername === null) {
    loginForm.classList.remove(HIDDEN_CLASSNAME);
    loginForm.addEventListener("submit", onLoginSubmit);
} else {
    greeting.innerText = `Hello ${saveUsername}`;
    greeting.classList.remove(HIDDEN_CLASSNAME);
    paintGreeting();
}*/