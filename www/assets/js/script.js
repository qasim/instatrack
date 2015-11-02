window.minTagID = ''

$(document).ready(function () {
  var tag = getTag()
  if (tag != '') {
    var grid = $('.grid')
    for (i = 0; i < 20; i++) {
      grid.append('<div class="item" id="index_' + i + '"></div>')
    }

    $("#tag").val(tag)
    getMedia(tag)
    var timer = setInterval(function () { getMedia(tag) }, 2000)
  }
})

function getTag() {
  var url = window.location.href
  if(url.indexOf('=') < 0) return ''
  return window.location.href.split('=')[1]
}

function getMedia(tag) {
  $.getJSON("/media?tag=" + tag + '&min_tag_id=' + window.minTagID, function (data) {
    if (data.data.length > 0) {
      window.minTagID = data.pagination.min_tag_id
    }
    var data = data.data
    var indices = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19]
    indices = shuffle(indices)
    for(i = 0; i < data.length; i++) {
      var item = $('#index_' + indices[i])
      item.css({
        'background-image': 'url(' + data[i].images.standard_resolution.url + ')'
      })
    }
  })
}

function shuffle(array) {
  var currentIndex = array.length, temporaryValue, randomIndex ;

  // While there remain elements to shuffle...
  while (0 !== currentIndex) {

    // Pick a remaining element...
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;

    // And swap it with the current element.
    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }

  return array;
}
