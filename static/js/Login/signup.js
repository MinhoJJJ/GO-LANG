$(document).ready(function () {

    var validateCk=true;
    var validateId=false;
    $('#checkDuplicateId').click(function () {
        //유효성 체크
        validateCheck("ID");
    });

    $('#signUpBtn').click(function () {
        var userId = $('#signupId').val();
        var password = $('#Password').val();
        var birthday = $('#Birthday').val();
        var email = $('#Email').val();
        var name = $('#Name').val();

        // var writeForm = $('#signupForm');
        // var other_data = writeForm.serializeArray();
        // var formData  = new FormData();

        // $.each(other_data,function(key,input){
        //     console.log(input.name+"= "+input.value);
        //     form.append(input.name, input.value);
        // });

        //유효성체크
        validateCheck("total");
        if(validateCk) {
            $.ajax({
                url: '/signUp.do', // 회원가입 로직
                method: 'POST',
                data: {
                    Id: userId, Password: password,
                    Birthday: birthday, Name: name, Email: email
                },
                success: function (response) {
                    if (response.result == "S") {
                        alert('회원가입에 성공하셨습니다.');
                        window.location.href = "/"; // 이동할 페이지 URL로 변경
                    } else {
                        alert('회원가입에 실패하셨습니다.');
                    }
                },
                error: function () {
                    alert('중복 확인 중 오류가 발생했습니다.');
                }
            });
        }
    });

    function validateEmail(email) {
        var isValid = /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
        // 이메일 형식 체크
        if (!isValid) {
            alert('이메일에 @가 빠졌습니다.');
            $('#Email').val("");
            validateCk=false;
        }
    }
    function validateCheck(ck) {
        var userId = $('#signupId').val();
        var password = $('#Password').val();
        var birthday = $('#Birthday').val();
        var email = $('#Email').val();
        var name = $('#Name').val();

        validateCk=true;
        if(ck=="ID"){
            validateId=true;
        }
        if(!validateId){
            alert('아이디 중복검사를 먼저 해주세요.');
            validateCk=false;
            return;
        }
        // 공란 체크
        if (userId === '') {
            alert('아이디를 입력하세요.');
            validateCk=false;
            return;
        }
        if (password === '' && ck!="ID") {
            alert('비밀번호를 입력하세요.');
            validateCk=false;
            return;
        }
        if (name === ''&& ck!="ID") {
            alert('이름를 입력하세요.');
            validateCk=false;
            return;
        }
        if (birthday === ''&& ck!="ID") {
            alert('생일을 입력하세요.');
            validateCk=false;
            return;
        }


        // 아이디 형식 체크
        var isValid = /^[a-z0-9]*$/.test(userId);
        if (!isValid) {
            // 유효하지 않은 문자가 포함된 경우 처리 (예: 경고 메시지 표시)
            alert('아이디는 영어 소문자와 숫자만 입력할 수 있습니다.');
            $('#signupId').val("");
            validateCk=false;
            return;
        }

        // AJAX 요청으로 아이디 중복 확인
        $.ajax({
            url: '/checkDuplicateIdAjax.do', // 실제 아이디 중복 확인 URL로 변경
            method: 'POST',
            async : false,
            data: { Id: userId },
            success: function (response) {
                if (response.result =="S") {
                    alert('이미 사용 중인 아이디입니다.');
                    $('#signupId').val("");
                    validateCk=false;
                } else {
                    if(ck=="ID"){
                        alert('사용 가능한 아이디입니다.');
                    }else{
                        var isValid = /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
                        // 이메일 형식 체크
                        if (!isValid) {
                            alert('이메일 형식이 틀렸습니다.');
                            $('#Email').val("");
                            validateCk=false;
                        }
                    }
                }
            },
            error: function () {
                alert('중복 확인 중 오류가 발생했습니다.');
            }
        });
    }
});