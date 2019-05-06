<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="<?= asset_link('css/bootstrap.min.css'); ?>">
    <script src="<?= asset_link('js/bootstrap.min.js') ?>"></script>
    <title>Login</title>
</head>
<body>
<div class="container">
    <h1>Welcome, visitor</h1>
    <h3>Please login</h3>
    <form method="post" action="<?= page_link('post-login') ?>">
        <div class="form-group">
            <label for="user">User</label>
            <input class="form-control" name="user">
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input class="form-control" name="password" type="password">
        </div>
        <button class="btn btn-info" type="submit">Login</button>
    </form>
</div>
</body>
</html>
