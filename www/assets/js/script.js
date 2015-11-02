window.minTagID = ''

$(document).ready(function () {
  var tag = getTag()
  if (tag != '') {
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
    console.log(data)
  })
}
