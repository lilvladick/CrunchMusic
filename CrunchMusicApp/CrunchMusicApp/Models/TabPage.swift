enum TabPage: String, CaseIterable {
    case music = "News"
    case playlist = "Breaking"
    case settings = "Settings"
    
    var icon: String {
        switch self {
        case .music: return "newspaper"
        case .playlist: return "exclamationmark.triangle"
        case .settings: return "gearshape"
        }
    }
    
    var label: String {
        self.rawValue
    }
}
