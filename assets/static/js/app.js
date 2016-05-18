$('#confirmAction').on('show.bs.modal', function (event) {

  var button = $(event.relatedTarget);
  var action = button.data('action');
 
  $(this).find('.modal-content form').on('submit', function(event) {

        $.ajax({
            method   : 'POST',
            url      : action,
            data     : $(this).serialize(),
            dataType : 'json',
            success  : function(data, event) {
                location.reload(true);
            },
            error: function(request, status, error) {
              
                switch (request.status) {
                    case 403:
                        alert("Authentication failed");
                        break;
                    default:
                        alert(error);
                        break;
                }
            }
        });

        return false
    });
});

