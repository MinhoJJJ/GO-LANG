$(document).ready(function () {

    // 수입, 지출 대분류 카테고리 로드
    loadMainCategories("MI");
    loadMainCategories("ME");
    loadMainCategories("ALL");

    // 수입 대분류 카테고리 추가
    $('#addIncomeMainCategoryButton').click(function () {
        var categoryName = $('#newIncomeMainCategory').val();
        var categoryColor = $('#newIncomeCategoryColor').val();
        if(categoryName=='' ||categoryName==null){
            alert("대분류 카테고리를 입력해주세요.");
            return;
        }
        insertCategory("MI",categoryName,"",categoryColor)
    });

    // 수입 소분류 카테고리 추가
    $('#addIncomeCategoryButton').click(function () {
        var mainCategoryName = $('#selectIncomeMainCategory').val();
        var categoryName = $('#newIncomeCategory').val();

        if(mainCategoryName=='' ||mainCategoryName==null){
            alert("대분류 카테고리를 선택해주세요.");
            return;
        }
       insertCategory("MI",mainCategoryName,categoryName,"");
    });

    // 지출 대분류 카테고리 추가
    $('#addExpenseMainCategoryButton').click(function () {
        var categoryName = $('#newExpenseMainCategory').val();
        var categoryColor = $('#newExpenseCategoryColor').val();
        if(categoryName=='' ||categoryName==null){
            alert("대분류 카테고리를 입력해주세요.");
            return;
        }
        insertCategory("ME",categoryName,"",categoryColor)
    });

    //카테고리 추가
    function insertCategory(main_type,main_CategoryName,sub_CategoryName,color){

        $.ajax({
            url: '/addCategory.do',
            method: 'POST',
            data: { id: "wat", main_type: main_type, color: color , main_CategoryName: main_CategoryName, sub_CategoryName: sub_CategoryName},
            success: function (response) {
                if (response.result == "S") {
                    alert(response.message);
                } else {
                    alert(response.message);
                }
                // 성공, 실패 유무에 관계없이 초기화
                $('#newIncomeMainCategory').val("");
                $('#newExpenseMainCategory').val("");

                setTimeout(function() {
                    location.reload();
                }, 10);
            },
            error: function () {
                alert('카테고리 추가에 실패하였습니다.');
            }
        });
    }
    // 수입, 지출 대분류 카테고리 로드
    function loadMainCategories(main_type) {
        $.ajax({
            url: '/loadMainCategories.do',
            method: 'POST',
            data: { id: "wat",main_type: main_type, },
            success: function(categories) {
                if(main_type=="MI"){
                    var select = $('#selectIncomeMainCategory');
                    select.empty();
                    select.append($('<option>').val('').text('대분류 선택'));
                    $.each(categories, function(i, category) {
                        select.append($('<option>').val(category.main_category).text(category.main_category));
                    });
                }else if(main_type=="ME"){
                    var select = $('#selectExpenseMainCategory');
                    select.empty();
                    select.append($('<option>').val('').text('대분류 선택'));
                    $.each(categories, function(i, category) {
                        select.append($('<option>').val(category.main_category).text(category.main_category));
                    });
                }else if(main_type=="ALL"){
                    var select = $('#selectMainCategory');
                    var select2 = $('#selectMainCategory2');

                    select.empty();
                    select2.empty();
                    select.append($('<option>').val('').text('대분류 선택'));
                    select2.append($('<option>').val('').text('대분류 선택'));

                    $.each(categories, function(i, category) {
                        select.append($('<option>').val(category.main_category).text(category.main_category));
                        select2.append($('<option>').val(category.main_category).text(category.main_category));
                    });
                }

            },
            error: function(xhr, status, error) {
                console.error('카테고리 로드 중 오류 발생:', error);
            }
        });
    }

});