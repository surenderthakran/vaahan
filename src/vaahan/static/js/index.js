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
  }
})();
