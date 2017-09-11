'use strict'

function getCurrentSimulation() {
  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/get_sim',
    {
      method: 'GET',
      headers: myHeaders,
      mode: 'cors',
      cache: 'default',
    }
  );

  fetch(myRequest)
  .then(function(response) {
    console.log(response);
    if(response.ok) {
      return response.json();
    } else {
      throw Error(response.statusText);
    }
  })
  .then(function(sim) {
    console.log(sim);
    _map.track = sim.track;
    _car.data = sim.car;
    initCanvas();
    drawMap();
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}

function updateRestartConf(restart) {
  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/update_restart_conf?restart=' + (restart?1:0),
    {
      method: 'GET',
      headers: myHeaders,
      mode: 'cors',
      cache: 'default',
    }
  );

  fetch(myRequest)
  .then(function(response) {
    console.log(response);
    if(response.ok) {
      return response;
    } else {
      throw Error(response.statusText);
    }
  })
  .then(function(res) {
    console.log(res);
  })
  .catch(function(err) {
    console.error(err);
  });
}

function getCar() {
  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/get_car',
    {
      method: 'GET',
      headers: myHeaders,
      mode: 'cors',
      cache: 'default',
    }
  );

  fetch(myRequest)
  .then(function(response) {
    console.log(response);
    if(response.ok) {
      return response.json();
    } else {
      throw Error(response.statusText);
    }
  })
  .then(function(car) {
    _car.data = car;
    console.log(_car.data);
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}

function initCar() {
  if (_map.track == undefined) {
    document.getElementById('status').textContent = "Select a track first!";
    return;
  }

  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/init_car',
    {
      method: 'GET',
      headers: myHeaders,
      mode: 'cors',
      cache: 'default',
    }
  );

  fetch(myRequest)
  .then(function(response) {
    console.log(response);
    if(response.ok) {
      return response.json();
    } else {
      throw Error(response.statusText);
    }
  })
  .then(function(car) {
    _car.data = car;
    console.log(_car.data);
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}

function driveCar() {
  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/drive_car',
    {
      method: 'GET',
      headers: myHeaders,
      mode: 'cors',
      cache: 'default',
    }
  );

  fetch(myRequest)
  .then(function(response) {
    console.log(response);
    if(response.ok) {
      return response.json();
    } else {
      throw Error(response.statusText);
    }
  })
  .then(function(car) {
    _car.data = car;
    console.log(_car.data);
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}

function stopCar() {
  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/stop_car',
    {
      method: 'GET',
      headers: myHeaders,
      mode: 'cors',
      cache: 'default',
    }
  );

  fetch(myRequest)
  .then(function(response) {
    console.log(response);
    if(response.ok) {
      return response.json();
    } else {
      throw Error(response.statusText);
    }
  })
  .then(function(car) {
    _car.data = car;
    console.log(_car.data);
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}
