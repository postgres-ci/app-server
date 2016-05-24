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

$('#changePassword, #userForm, #projectForm').on('hide.bs.modal', function (event) {
    $(this).find('.modal-dialog form').trigger("reset");
});

$('#projectForm').on('show.bs.modal', function (event) {

    var button = $(event.relatedTarget);

    $(this).find('.modal-dialog form').attr('action', button.data('action'));
    $(this).find('.modal-content .modal-title').html(button.data('title'));

    $.ajax({
        method   : 'GET',
        url      : button.data('source'),
        data     : $(this).serialize(),
        dataType : 'json',
        success  : function(data, event) {
            $('#project_name').val(data.project_name);
            $('#repository_url').val(data.repository_url);
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
});

$('#resetPasswordForm').on('show.bs.modal', function (event) {

    var button = $(event.relatedTarget);

    $(this).find('.modal-dialog form').attr('action', button.data('action'));
});

$('#updateUserForm').on('show.bs.modal', function (event) {

    var button = $(event.relatedTarget);

    $(this).find('.modal-dialog form').attr('action', button.data('action'));

    $.ajax({
        method   : 'GET',
        url      : button.data('source'),
        data     : $(this).serialize(),
        dataType : 'json',
        success  : function(data, event) {
            $('#user_name').val(data.user_name);
            $('#user_email').val(data.user_email);
            $('#is_superuser').attr('checked', data.is_superuser);
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
});

$('#changePassword form, #projectForm form, #addUserForm form, #updateUserForm form, #resetPasswordForm form').bootstrap3Validate(function(e, data) {

    e.preventDefault();

    $.ajax({
        method   : 'POST',
        url      : $(this).attr('action'),
        data     : $(this).serialize(),
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
