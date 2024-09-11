const shortUrl = document.querySelector("#short-url");

const form = document.querySelector("#form");

const copy = document.querySelector("#copy");

form.addEventListener("submit", (e) => {
  e.preventDefault();

  const url = document.querySelector("#url").value;

  if (url === "") {
    window.alert("please enter a url");
    return;
  }

  const data = {
    long_url: url,
  };

  fetch("http://localhost:8080/short", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((data) => data.json())
    .then((url) => {
      shortUrl.textContent = url?.short_url;
      copy.style.display = "inline";
    });
});

copy.addEventListener("click", (e) => {
  navigator.clipboard
    .writeText(shortUrl.textContent)
    .then(function () {
      document.getElementById("copy").textContent = "Copied to clipboard!";
    })
    .catch(function (error) {
      document.getElementById("copy").textContent = "Failed to copy text!";
      console.error("Error copying text: ", error);
    });
});
