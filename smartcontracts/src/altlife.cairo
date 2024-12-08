use starknet::{ContractAddress};

#[starknet::interface]
trait IWalletMapping<TContractState> {
    fn set_wallet_data_for(ref self: TContractState, wallet: ContractAddress, data: felt252);
    fn get_wallet_data(self: @TContractState, wallet: ContractAddress) -> felt252;
}

#[starknet::contract]
mod WalletMapping {
    use starknet::{
        ContractAddress
    };
     use starknet::storage::{
        Map, StorageMapReadAccess, StorageMapWriteAccess
    };

    // Storage declaration
    #[storage]
    struct Storage {
        alt_life_data: Map::<ContractAddress, felt252>,
    }

    // Contract implementation
    #[abi(embed_v0)]
    impl WalletMapping of super::IWalletMapping<ContractState> {
        fn set_wallet_data_for(ref self: ContractState, wallet: ContractAddress, data: felt252) {            
            // Write the data to storage
            self.alt_life_data.write(wallet, data);
        }

        fn get_wallet_data(self: @ContractState, wallet: ContractAddress) -> felt252 {            
            // Read and return the data from storage
            self.alt_life_data.read(wallet)
        }
    }
}