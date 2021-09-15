// NavIgation Menu JS Code
let body = document.querySelector("body");
let navBar = document.querySelector(".nav");
let menuBtn = document.querySelector(".navbar-toggler-icon");
let cancelBtn = document.querySelector(".cancel-btn");

menuBtn.onclick = function(){
  menuBtn.style.opacity = "0";
  menuBtn.style.pointerEvents = "none";
  cancelBtn.style.opacity = "1";
  cancelBtn.style.display = "block";
}
cancelBtn.onclick = function(){
  cancelBtn.style.opacity = "0";
  menuBtn.style.opacity = "1";
  menuBtn.style.pointerEvents = "auto";
}

// Navigation Bar Close While We Click On Navigation Links
let navLinks = document.querySelectorAll(".navbar-nav li a");
for (var i = 0; i < navLinks.length; i++) {
  navLinks[i].addEventListener("click" , function() {
    navBar.classList.remove("show");
    menuBtn.style.opacity = "1";
    cancelBtn.style.opacity = "0";
    menuBtn.style.pointerEvents = "auto";
  });
}
