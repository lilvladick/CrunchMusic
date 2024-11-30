import Foundation

struct Comments: Codable, Identifiable {
    var id: Int
    var newsId: Int
    var authorId: Int
    var commentContent: String
    var createdAt: String
    var updatedAt: String
    
    enum CodingKeys: String, CodingKey {
        case id
        case newsId = "news_id"
        case authorId = "author_id"
        case commentContent = "comment_content"
        case createdAt = "created_at"
        case updatedAt = "updated_at"
    }
}
