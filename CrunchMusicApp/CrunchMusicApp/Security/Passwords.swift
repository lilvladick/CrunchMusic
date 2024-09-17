import Foundation
import Security

public extension SecureStorage {
    func addPassword(_ password: String, for account: String) {
        var query: [CFString: Any] = [:]
        query[kSecClass] = kSecClassGenericPassword
        query[kSecAttrAccount] = account
        query[kSecValueData] = password.data(using: .utf8)
        
        do {
            try addItem(query: query)
        } catch {
            return
        }
    }
    
    func getPassword(for account: String) -> String? {
            var query: [CFString: Any] = [:]
            query[kSecClass] = kSecClassGenericPassword
            query[kSecAttrAccount] = account
            
            var result: [CFString: Any]?
            
            do {
                result = try findItem(query: query)
            } catch {
                return nil
            }
            
            if let data = result?[kSecValueData] as? Data {
                return String(data: data, encoding: .utf8)
            } else {
                return nil
            }
        }
    
    func updatePassword(_ password: String, for account: String) {
        guard let _ = getPassword(for: account) else {
            addPassword(password, for: account)
            return
        }
        
        var query: [CFString: Any] = [:]
        
        query[kSecClass] = kSecClassGenericPassword
        query[kSecAttrAccount] = account
        
        var attridutesToUpdate: [CFString: Any] = [:]
        attridutesToUpdate[kSecValueData] = password.data(using: .utf8)
        
        do {
            try updateItem(query: query, attributesToUpdate: attridutesToUpdate)
        } catch {
            return
        }
    }
    
    func deletePassword(for account: String) {
        var query: [CFString: Any] = [:]
        query[kSecClass] = kSecClassGenericPassword
        query[kSecAttrAccount] = account
        
        do {
            try deleteItem(query: query)
        } catch {
            return
        }
    }
}
