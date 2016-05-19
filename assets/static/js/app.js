$('#confirmAction').on('show.bs.modal', function (event) {

    $(this).find('.modal-content form').attr('action', $(event.relatedTarget).data('action'));
    
}).on('hide.bs.modal', function (event) {

    $(this).find('.modal-content form input[type=password]').val('');
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
                case 401:
                    alert("Authentication failed");
                    break;
                case 403:
                    alert("Access denied");
                    break;
                default:
                    alert(error);
                    break;
            }
        }
    });

    return false
});
