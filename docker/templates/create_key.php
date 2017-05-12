<?php

$b = random_bytes(54);

$file_name = "super-secret-key.dat";

$f = fopen(dirname(__DIR__).'\settings\\'.$file_name, 'w+');
fwrite($f, $b."\0");
