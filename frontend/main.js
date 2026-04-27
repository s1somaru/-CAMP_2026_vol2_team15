const loginButton = document.getElementById('loginButton');
const emailInput = document.getElementById('emailInput');
const passwordInput = document.getElementById('passwordInput');
const message = document.getElementById('message');

loginButton.addEventListener('click', function() {
    // 入力欄が空っぽでないかチェックする
    if (emailInput.value === '' || passwordInput.value === '') {
        message.textContent = 'メールアドレスとパスワードを入力してください。';
        message.style.color = '#e53935'; // エラー時は赤色
    } else {
        // 成功した場合は、home.htmlへ画面を移動させる
        window.location.href = 'home.html';
    }
});