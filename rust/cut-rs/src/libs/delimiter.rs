pub fn get_delimiter(options: &[String]) -> String {
    if let Some(delimiter_option) = options.iter().find(|&opt| opt.starts_with("-d")) {
        if let Some(delimiter) = delimiter_option.strip_prefix("-d") {
            return delimiter.to_string();
        }
    }

    "".to_string()
}
