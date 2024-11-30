enum TabPage: String, CaseIterable {
    case news = "Новости"
    case breakingNews = "Срочные"
    case settings = "Settings"
    
    var icon: String {
        switch self {
        case .news: return "newspaper"
        case .breakingNews: return "exclamationmark.triangle"
        case .settings: return "gearshape"
        }
    }
    
    var label: String {
        self.rawValue
    }
}
