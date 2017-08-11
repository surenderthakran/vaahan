'use strict';

(function() {
  document.addEventListener('DOMContentLoaded', function(){
    var myHeaders = new Headers();
    var myRequest = new Request(window.location.origin + '/api/get_track', {
      method: 'GET',
      headers: myHeaders,
      mode: 'cors',
      cache: 'default',
    });

    fetch(myRequest)
    .then(function(response) {
      console.log(response);
      if(response.ok) {
        return response.json();
      } else {
        throw Error(response.statusText);
      }
    })
    .then(function(track) {
      drawTrack(track);
    })
    .catch(function(err) {
      console.error(err);
    });
  });

  function drawTrack(trackData) {
    console.log("inside drawTrack()");
    console.log(trackData);
    let trackCanvas = document.getElementById('track');
    trackCanvas.height = trackData.height;
    trackCanvas.width = trackData.width;
    var ctx = trackCanvas.getContext("2d");
    drawRoad(ctx, trackData.road);
    drawStartFinishLine(ctx, trackData.startingLine, trackData.finishingLine);
  }

  function drawRoad(ctx, roadData) {
    ctx.beginPath();
    for (let i = 0, len = roadData.length; i < len; i++) {
      let line = roadData[i];
      if (i === 0) {
        ctx.moveTo(line.startX, line.startY);
      } else {
        ctx.lineTo(line.startX, line.startY);
      }
      ctx.lineTo(line.endX, line.endY);
    }
    ctx.closePath();
    ctx.stroke();
    ctx.fillStyle = "#777777";
    ctx.fill();
  }

  function drawStartFinishLine(ctx, startingLine, finishingLine) {
    ctx.beginPath();
    ctx.moveTo(startingLine.startX, startingLine.startY);
    ctx.lineTo(startingLine.endX, startingLine.endY);
    ctx.moveTo(finishingLine.startX, finishingLine.startY);
    ctx.lineTo(finishingLine.endX, finishingLine.endY);
    ctx.closePath();
    ctx.strokeStyle="#ffffff";
    ctx.LineWidth = 150;
    ctx.stroke();
  }
})();
