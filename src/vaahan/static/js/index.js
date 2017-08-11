'use strict';

(function() {
  document.addEventListener('DOMContentLoaded', function(){
    console.log("loaded");
    console.log(window.location.origin);

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
      console.log(track);
    })
    .catch(function(err) {
      console.error(err);
    });
  });
})();
