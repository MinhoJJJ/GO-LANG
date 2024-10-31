$(document).ready(function() {
    $('#expenseBtn').click(function() {
        $(this).addClass('active');
        $('#incomeBtn').removeClass('active');

        // 결제 종류 선택 요소 보이기
        $('.payment-method').show();
        // 카테고리 선택 요소 원래대로 변경
        $('#category').empty();
        $('#category').append('<option value="food">식비</option>');
        $('#category').append('<option value="transport">교통비</option>');
        $('#category').append('<option value="entertainment">엔터테인먼트</option>');
        $('#category').append('<option value="other">기타</option>');
    });

    $('#incomeBtn').click(function() {
        $(this).addClass('active');
        $('#expenseBtn').removeClass('active');

        // 결제 종류 선택 요소 숨기기
        $('.payment-method').hide();
        // 카테고리 선택 요소 변경
        $('#category').empty();
        $('#category').append('<option value="salary">월급</option>');
        $('#category').append('<option value="allowance">용돈</option>');
    });
});