import Foundation

class NetworkManager {
    
    static func fetchData<T: Codable>(from urlString: String) async throws -> T {
        guard let url = URL(string: urlString) else {
            throw NetworkError.invalidURL
        }

        do {
            let (data, response) = try await URLSession.shared.data(from: url)
            
            if let httpResponse = response as? HTTPURLResponse {
                if httpResponse.statusCode != 200 {
                    throw NetworkError.invalidResponse
                }
            } else {
                throw NetworkError.invalidResponse
            }

            if data.isEmpty {
                throw NetworkError.noData
            }

            let decodedData = try JSONDecoder().decode(T.self, from: data)
            return decodedData
        } catch {
            throw error
        }
    }
}


enum NetworkError: Error {
    case invalidURL
    case invalidResponse
    case noData
}
