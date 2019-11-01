function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

async function initMap() {
  // starting coords
  let lastLat = 42.47012;
  let lastLng = -83.05653;
  
  // initialize map and marker
  let initPos = {lat: lastLat, lng: lastLng}
  let map = new google.maps.Map(document.getElementById('map'), {
    zoom: 20,
    center: initPos,
    disableDefaultUI: true,
    mapTypeId: 'satellite',
  }); 
  let marker = new google.maps.Marker({
    map: map,
    animation: google.maps.Animation.DROP,
    position: initPos,
  });

  // poll rest api at /position
  // update marker and pan if needed
  while(true) {
    let response = await fetch(window.location.origin + "/position");
    let data = await response.json();
    let lat = data.lat;
    let lng = data.lng;
    console.log("polled, at (" + lat + ", " + lng + ")");
    if((lat != lastLat) || (lng != lastLng)) {
      console.log("moved to: (" + lat + ", " + lng + ")");
      pos = new google.maps.LatLng(lat, lng);
      marker.setPosition(pos);
      map.panTo(pos);
      lastLat = lat;
      lastLng = lng;
    }
    await sleep(1000);
  }
}
