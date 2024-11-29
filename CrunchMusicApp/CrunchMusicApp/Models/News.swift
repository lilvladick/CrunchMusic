import Foundation

struct News: Codable, Identifiable {
    let id: Int
    let title: String
    let newsContent: String
    let authorID: Int?
    let categoryID: Int?
    let isBreaking: Bool
    let publishedAt: String
    let updatedAt: String

    enum CodingKeys: String, CodingKey {
        case id
        case title
        case newsContent = "news_content" 
        case authorID = "author_id"
        case categoryID = "category_id"
        case isBreaking = "is_breaking"
        case publishedAt = "published_at"
        case updatedAt = "updated_at"
    }
}

