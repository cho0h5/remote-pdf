var loadingTask = pdfjsLib.getDocument("example.pdf");

function renderPage(n) {
  loadingTask.promise.then(function (pdf) {
    pdf.getPage(n).then(function (page) {
      var scale = 1;
      var viewport = page.getViewport({ scale: scale });

      var canvas = document.getElementById("the-canvas");
      var context = canvas.getContext("2d");
      canvas.height = viewport.height;
      canvas.width = viewport.width;
      canvas.font = ""

      var renderContext = {
        canvasContext: context,
        viewport: viewport,
      };
      page.render(renderContext);
    });
  });
}
