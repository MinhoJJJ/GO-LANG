$(document).ready(function () {
    $('#checkDuplicateId').click(function () {
        const userId = $('#signupId').val().trim();
        if (userId === '') {
            alert('아이디를 입력하세요.');
            return;
        }

        // AJAX 요청으로 아이디 중복 확인 (예시)
        $.ajax({
            url: '/check-duplicate-id', // 실제 아이디 중복 확인 URL로 변경
            method: 'POST',
            data: { userId: userId },
            success: function (response) {
                if (response.isDuplicate) {
                    alert('이미 사용 중인 아이디입니다.');
                } else {
                    alert('사용 가능한 아이디입니다.');
                }
            },
            error: function () {
                alert('중복 확인 중 오류가 발생했습니다.');
            }
        });
    });

    $('#signupForm').submit(function (e) {
        const email = $('#signupEmail').val().trim();
        if (!validateEmail(email)) {
            alert('유효한 이메일 주소를 입력하세요.');
            e.preventDefault();
        }
    });

    function validateEmail(email) {
        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return re.test(email);
    }
});