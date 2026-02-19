(* Logic to determine if a number is even *)
let is_even n =
  n mod 2 = 0

(* Recursive function to sum even numbers from 1 to current n *)
let rec sum_nums n =
  if n <= 0 then
    0
  else if is_even n then
    n + sum_nums (n - 1)
  else
    sum_nums (n - 1)

(* Main execution block for input handling *)
let () =
  print_string "Enter N: ";
  let n = read_int () in
  let result = sum_nums n in
  Printf.printf "The sum of even numbers from 1 to %d is: %d\n" n result