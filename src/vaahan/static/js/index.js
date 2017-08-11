'use strict';

(function() {
  document.addEventListener('DOMContentLoaded', function(){
    console.log("loaded");
    console.log(window.location.origin);

    var myHeaders = new Headers();
    var myRequest = new Request(window.location.origin + '/api/get_map', {
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
    .then(function(map) {
      drawMap(map);
    })
    .catch(function(err) {
      console.error(err);
    });
  });

  function drawMap(mapData) {
    console.log("inside drawMap()");
    console.log(mapData);
    let mapCanvas = document.getElementById('map');
    mapCanvas.height = mapData.height;
    mapCanvas.width = mapData.width;
    var ctx = mapCanvas.getContext("2d");
    for (let i = 0, len = mapData.road.length; i < len; i++) {
      console.log(mapData.road[i]);
      let line = mapData.road[i]
      ctx.moveTo(line.startX, line.startY);
      ctx.lineTo(line.endX, line.endY);
      ctx.stroke();
    }
  }
})();
