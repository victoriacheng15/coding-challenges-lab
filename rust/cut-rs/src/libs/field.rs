pub fn get_field_numbers(options: &[String]) -> Vec<u32> {
    let has_field = options.iter().any(|opt| opt.starts_with("-f"));
    let mut field_number = Vec::new();
    if has_field {
        let field = options
            .iter()
            .find(|opt| opt.starts_with("-f"))
            .unwrap()
            .split("f")
            .nth(1)
            .unwrap();

        for num in field.chars() {
            if num.is_digit(10) {
                field_number.push(num.to_digit(10).unwrap());
            }
        }
    }

    field_number
}
