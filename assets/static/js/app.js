$('#confirmAction').on('show.bs.modal', function (event) {

    var button = $(event.relatedTarget);
 
    $(this).find('.modal-content form').attr('action', button.data('action'));
});

$('#confirmAction').find('.modal-content form').on('submit', function(event) {

    $.ajax({
        method   : 'POST',
        url      : $(this).attr('action'),
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