#[cfg(test)]
mod tests {
    use std::str::FromStr;

    use ethers::types::{I256, U256};

    #[test]
    fn add_1() {
        assert_eq!(
            U256::from_str("10")
                .unwrap()
                .overflowing_add(U256::from_str("10").unwrap())
                .0,
            U256::from_str("20").unwrap()
        );
    }

    #[test]
    fn add_2() {
        assert_eq!(
            U256::from_str_radix(
                "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
                16
            )
            .unwrap()
            .overflowing_add(U256::from_str_radix("0x1", 16).unwrap())
            .0,
            U256::from_str_radix("0x0", 16).unwrap()
        );
    }

    #[test]
    fn mul_1() {
        assert_eq!(
            I256::from_str("10").unwrap() * I256::from_str("10").unwrap(),
            I256::from_str("100").unwrap()
        );
    }

    #[test]
    fn mul_2() {
        assert_eq!(
            U256::from_str_radix(
                "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
                16,
            )
            .unwrap()
            .overflowing_mul(U256::from_str_radix("0x2", 16).unwrap())
            .0,
            U256::from_str_radix(
                "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE",
                16,
            )
            .unwrap()
        );
    }

    #[test]
    fn sub_1() {
        assert_eq!(
            I256::from_str("10")
                .unwrap()
                .overflowing_sub(I256::from_str("10").unwrap())
                .0,
            I256::from_str("0").unwrap()
        );
    }

    #[test]
    fn sub_2() {
        assert_eq!(
            U256::from_str_radix("0x0", 16,)
                .unwrap()
                .overflowing_sub(U256::from_str_radix("0x1", 16).unwrap())
                .0,
            U256::from_str_radix(
                "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
                16,
            )
            .unwrap()
        );
    }

    #[test]
    fn div_1() {
        assert_eq!(
            I256::from_str("10")
                .unwrap()
                .overflowing_div(I256::from_str("10").unwrap())
                .0,
            I256::from_str("1").unwrap()
        );
    }

    #[test]
    fn div_2() {
        assert_eq!(
            I256::from_str("1")
                .unwrap()
                .overflowing_div(I256::from_str("2").unwrap())
                .0,
            I256::from_str("0").unwrap()
        );
    }

    #[test]
    fn sdiv_1() {
        assert_eq!(
            I256::from_str("10")
                .unwrap()
                .overflowing_div_euclid(I256::from_str("10").unwrap())
                .0,
            I256::from_str("1").unwrap()
        );
    }

    #[test]
    fn sdiv_2() {
        assert_eq!(
            I256::from_hex_str(
                "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE",
            )
            .unwrap()
            .overflowing_div_euclid(
                I256::from_hex_str(
                    "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
                )
                .unwrap()
            )
            .0,
            I256::from_str("2").unwrap()
        );
    }
}
