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

$('#changePassword').on('hide.bs.modal', function (event) {
    $(this).find('.modal-dialog form').trigger("reset");
});

$('#changePassword form').bootstrap3Validate(function(e, data) {

    e.preventDefault();

    $.ajax({
        method   : 'POST',
        url      : '/password/change/',
        data     : data,
        dataType : 'json',
        success  : function(data, event) {
            location.reload(true);
        },
        error: function(request, status, error) {

            var data = JSON.parse(request.responseText);

            if (data.error !== undefined) {

                alert(data.error);

                return;
            } 
          
            alert(error);
        }
    });

    return false
});

$('#addProject form').bootstrap3Validate(function(e, data) {

    e.preventDefault();

    $.ajax({
        method   : 'POST',
        url      : '/project/add/',
        data     : data,
        dataType : 'json',
        success  : function(data, event) {
            location.reload(true);
        },
        error: function(request, status, error) {

            var data = JSON.parse(request.responseText);

            if (data.error !== undefined) {

                alert(data.error);

                return;
            } 
          
            alert(error);
        }
    });

    return false
});

