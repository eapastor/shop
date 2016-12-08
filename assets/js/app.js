$(document).ready(function() {
    $('#list').click(
        function(event){
            event.preventDefault();
            $('#products .item').removeClass('grid-group-item');
            $('#products .item').addClass('list-group-item');
    });
    $('#grid').click(
        function(event){
            event.preventDefault();
            $('#products .item').removeClass('list-group-item');
            $('#products .item').addClass('grid-group-item');
        });
    
    $('.inCart-btn').click( function (event) {

        console.log('inCart')

        if ($(this).attr('inCart') == 0){
            $(this).attr('inCart', 1);
            $(this).children('.inCart-btn-text')[0].innerText = '                                            В КОРЗИНЕ                                        '
            $(this).toggleClass('btn-success', 'btn-default')
            console.log('id:', $(this).children('.item_id').val(), 'amp:', $(this).children('.amp').val())
            var jqxhr = $.post( '/cart', { id: $(this).children('.item_id').val(), amp: $(this).children('.amp').val() }, function(data) {
                console.log( 'success' );
            })
                .done(function(data) {
                    console.log('ok', data.amp)
                })
                .fail(function(data) {
                    console.log('false', data)
                })
            } else {
                location.href = '/cart';
            }
        }
    )

    $('.amp').change( function (event) {
        console.log( $(this).siblings('.inCart-btn').get())
        $(this).siblings('.inCart-btn').css("background", "red");
        //.toggleClass('btn-success', 'btn-default')
        this.value=this.value.replace(/[^0-9]+/g,'');
        if ($(this).val()>25) $(this).val(25);
        if ($(this).val()<1) $(this).val(1);
    })
});