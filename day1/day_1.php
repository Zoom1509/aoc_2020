<?php
$file_1 = fopen("day_1_input.txt", "r");
$file_2 = fopen("day_1_input.txt", "r");
$file_3 = fopen("day_1_input.txt", "r");


while(!feof($file_1)) {

    $line_1 = intval(fgets($file_1));

    while(!feof($file_2)) {
        $line_2 = intval(fgets($file_2));

        while(!feof($file_3)) {
            $line_3 = intval(fgets($file_3));
            $sum = $line_1 + $line_2 + $line_3;
            if($sum == 2020) {
                $result = $line_1 * $line_2 * $line_3;
                echo $result;
                exit();
            }
        }
        rewind($file_3);
        
    }
    rewind($file_2);
}
?>