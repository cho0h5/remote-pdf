var currentPage = 1;

setInterval(function () {
  fetch("http://cho0h5.iptime.org:8080/status")
    .then((data) => data.json())
    .then((data) => data["PageNumber"])
    .then((page) => {
      if (currentPage != page) {
        currentPage = page;
        renderPage(page);
      }
    });
}, 50);
