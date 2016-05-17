$('#confirmAction').on('show.bs.modal', function (event) {
  var button = $(event.relatedTarget);
  var action = button.data('action');


  $(this).find('.modal-body form').attr('action', action);
})