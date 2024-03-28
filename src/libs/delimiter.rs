pub fn get_delimiter(options: &[String]) -> String {
    let _has_delimiter = options.iter().any(|opt| opt.starts_with("-d"));

    //  wait to extract the delimiter for later

    return ",".to_string();
}
