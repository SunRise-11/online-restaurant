$( document ).ready(function() {
    $( "form" ).on( "submit", function(e) {
           var dataString = $(this).serialize();
          //  alert(dataString); return false;

           $.ajax({
              type: 'POST',
              url: '/order',
              data: dataString,
              success: function (response) {
                  $('.toast').toast('show');
              }
       });

      e.preventDefault();
    });
        

  });
